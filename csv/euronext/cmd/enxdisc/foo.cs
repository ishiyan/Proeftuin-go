        internal static void UpdateTask(string downloadPath)
        {
            EuronextActualInstruments.DownloadOverwriteExisting = Properties.Settings.Default.DownloadOverwriteExisting;
            EuronextActualInstruments.DownloadRetries = Properties.Settings.Default.DownloadRetries;
            EuronextActualInstruments.DownloadTimeout = Properties.Settings.Default.DownloadTimeout;

            DateTime dateTime = DateTime.Now;
            Dictionary<string, EuronextActualInstruments.InstrumentInfo> actualInstrumentsDictionary =
                EuronextActualInstruments.Fetch(downloadPath, userAgent: Properties.Settings.Default.UserAgent);

            foreach (var v in EuronextActualInstruments.UnknownMicDictionary)
            {
                Trace.TraceError("Unknown MIC: {0}", v.Key);
                foreach (var isinInfo in actualInstrumentsDictionary)
                {
                    if (isinInfo.Value.Mic == v.Key)
                        Trace.TraceError("------------ {0}, {1}", isinInfo.Key, isinInfo.Value.Type);
                }
            }

            BackupXmlFile(Properties.Settings.Default.MainEuronextIndexPath, dateTime);
            BackupXmlFile(Properties.Settings.Default.OtherEuronextIndexPath, dateTime);
            /* BackupXmlFile(Properties.Settings.Default.UninterestedEuronextIndexPath, dateTime); */

            XDocument xdocMain = XDocument.Load(Properties.Settings.Default.MainEuronextIndexPath
                /* , LoadOptions.PreserveWhitespace | LoadOptions.SetLineInfo */);
            List<XElement> xelistMain = xdocMain.XPathSelectElements("/instruments/instrument").ToList();

            XDocument xdocUninterested = XDocument.Load(Properties.Settings.Default.UninterestedEuronextIndexPath
                /* , LoadOptions.PreserveWhitespace | LoadOptions.SetLineInfo */);
            xelistMain.AddRange(xdocUninterested.XPathSelectElements("/instruments/instrument").ToList());

            XDocument xdocOther = XDocument.Load(Properties.Settings.Default.OtherEuronextIndexPath
                /*, LoadOptions.PreserveWhitespace | LoadOptions.SetLineInfo*/);
            List<XElement> xelistOther = xdocOther.XPathSelectElements("/instruments/instrument").ToList();

            xdocMain.XPathSelectElement("/instruments").Add(new XComment(dateTime.ToString(" yyyyMMdd_HHmmss ")));
            xdocOther.XPathSelectElement("/instruments").Add(new XComment(dateTime.ToString(" yyyyMMdd_HHmmss ")));

            string[] micListMain = Properties.Settings.Default.MainEuronextMics.Split(';');
            foreach (var kvp in actualInstrumentsDictionary)
            {
                var ii = kvp.Value;
                List<XElement> matchListMain = xelistMain.FindAll(xel => xel.MatchesIsinMicSymbol(ii));
                ii.IsApproved = matchListMain.Count > 0;
                if (matchListMain.Count > 1)
                {
                    Trace.TraceError("Approved: {0} duplicate matches for \"{1}\":", matchListMain.Count, ii.Key);
                    int i = 0;
                    foreach (var xElement in matchListMain)
                    {
                        Trace.TraceError("{0}: element [{1}]", ++i, xElement.ToString(SaveOptions.None));
                    }
                }
                List<XElement> matchListOther = xelistOther.FindAll(xel => xel.MatchesIsinMicSymbol(ii));
                ii.IsDiscovered = matchListOther.Count > 0;
                if (matchListOther.Count > 1)
                {
                    Trace.TraceError("Discovered: {0} duplicate matches for \"{1}\":", matchListOther.Count, ii.Key);
                    int i = 0;
                    foreach (var xElement in matchListOther)
                    {
                        Trace.TraceError("{0}: element [{1}]", ++i, xElement.ToString(SaveOptions.None));
                    }
                }
                if (Properties.Settings.Default.Enrich && !ii.IsApproved && !ii.IsDiscovered)
                {
                    Trace.TraceInformation("Discovered {0}: \"{1}\":", ii.Type, ii.Key);
                    XElement xel = micListMain.Contains(ii.Mic) ? xdocMain.XPathSelectElement("/instruments") : xdocOther.XPathSelectElement("/instruments");
                    XElement xelNew = ii.NewInstrumentElement(true);
                    try
                    {
                    switch (ii.Type)
                    {
                        case EuronextInstrumentXml.Index:
                            xelNew.EnrichIndexElement(Properties.Settings.Default.UserAgent);
                            break;
                        case EuronextInstrumentXml.Stock:
                            xelNew.EnrichStockElement(Properties.Settings.Default.UserAgent);
                            break;
                        case EuronextInstrumentXml.Etv:
                            xelNew.EnrichEtvElement(Properties.Settings.Default.UserAgent);
                            break;
                        case EuronextInstrumentXml.Etf:
                            xelNew.EnrichEtfElement(Properties.Settings.Default.UserAgent);
                            XElement xelInav = xelNew.InavElementFromEtf();
                            if (null != xelInav)
                            {
                                xelInav.EnrichInavElement(Properties.Settings.Default.UserAgent);
                                if (xel != null) xel.Add(xelInav);
                            }
                            break;
                        case EuronextInstrumentXml.Fund:
                            xelNew.EnrichFundElement(Properties.Settings.Default.UserAgent);
                            break;
                        default:
                            xelNew.EnrichSearchInstrument(userAgent: Properties.Settings.Default.UserAgent);
                            break;
                    }
                    }
                    catch (Exception ex)
                    {
                        Trace.TraceError("Failed to enrich {0}: \"{1}\": {2}", ii.Type, ii.Key, ex);
                    }

                    if (xel != null) xel.Add(xelNew);
                }
            }
            try
            {
                xdocMain.Save(Properties.Settings.Default.MainEuronextIndexPath, SaveOptions.None);
            }
            catch (Exception ex)
            {
                Trace.TraceError("Failed to save {0}: {1}", Properties.Settings.Default.MainEuronextIndexPath, ex);
            }
            try
            {
                xdocOther.Save(Properties.Settings.Default.OtherEuronextIndexPath, SaveOptions.None);
            }
            catch (Exception ex)
            {
                Trace.TraceError("Failed to save {0}: {1}", Properties.Settings.Default.MainEuronextIndexPath, ex);
            }
        }
