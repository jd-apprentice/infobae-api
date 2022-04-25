interface TextParse {
  _text: string;
}

interface DataParse {
  slice(
    arg0: number,
    size:
      | number
      | import("qs").ParsedQs
      | number[]
      | import("qs").ParsedQs[]
      | undefined
  ): DataParse[];
  lastmod: TextParse;
  loc: TextParse;
}

interface JsonParsedProps {
  _declaration: {
    _attributes: {
      version: string;
      encoding: string;
    };
  };
  _text: string;
  urlset: {
    _attributes: {
      xmlns?: string;
      "xmlns:xsi"?: string;
      "xsi:schemaLocation"?: string;
      "xmlns:xhtml"?: string;
      "xmlns:atom"?: string;
      "xmlns:dc"?: string;
      "xmlns:media"?: string;
      "xmlns:geo"?: string;
      "xmlns:opensearch"?: string;
      "xmlns:rss"?: string;
      "xmlns:rdf"?: string;
      "xmlns:thr"?: string;
      "xmlns:xsd"?: string;
      "xmlns:content"?: string;
      "xmlns:itunes"?: string;
      "xmlns:atom10"?: string;
    };
    url: Array<DataParse>;
  };
}

interface SizeProps {
  size: number;
}

type ParsedQs = import("qs").ParsedQs;

export { DataParse, JsonParsedProps, ParsedQs, SizeProps };
