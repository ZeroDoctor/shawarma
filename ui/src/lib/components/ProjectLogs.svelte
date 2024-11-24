<script>
	import { onDestroy, onMount } from "svelte";
    import hljs from 'highlight.js/lib/core';
    import { genericlog } from "$lib/highlight.log";
    import { CodeBlock, storeHighlightJs } from "@skeletonlabs/skeleton";

    hljs.registerLanguage('genericlog', genericlog);
    storeHighlightJs.set(hljs);

    export let projectId = '';

    onMount(() => {
        console.log('mount project logs =>', projectId)
    });

    onDestroy(() => {
        console.log('destory project logs =>', projectId)
    });
</script>

<CodeBlock language="genericlog" lineNumbers={true} code={`
INFO: This is an info log
DEBUG: This is a debug log
ERROR: This is an error log
WARN: This is a warning log
TRACE: This is a trace log
CRITICAL: This is a critical log
20.164.151.111 - INFO - [20/Aug/2015:22:20:18 -0400] "GET /mywebpage/index.php HTTP/1.1" 403 772 "-" "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_6_8) AppleWebKit/535.1 (KHTML, like Gecko) Chrome/13.0.782.220 Safari/535.1"
`}>
</CodeBlock>
