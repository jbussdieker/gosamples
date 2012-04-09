#include <stdio.h>
#include <libxml/parser.h>
#include <libxml/tree.h>

static const char *document = "<doc/>";

void DoIt() {
	LIBXML_TEST_VERSION

	xmlDocPtr doc;
	doc = xmlReadMemory(document, sizeof(document), "noname.xml", NULL, 0);
	if (doc == NULL) {
        fprintf(stderr, "Failed to parse document\n");
		return;
    } else {
		int size = 0;
		xmlChar *mem;
		xmlDocDumpMemory(doc, &mem, &size);
		fprintf(stdout, "Document size: %d\nDocument: \n%s", size, mem);
		xmlFree(mem);
	}
    xmlFreeDoc(doc);
	xmlCleanupParser();
	xmlMemoryDump();
}
