internal static void UpdateTask(string downloadPath)
{
    // some code ...
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
    // Rest oif code ...
}
