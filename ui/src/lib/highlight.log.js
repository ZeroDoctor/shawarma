
/**
 * @param {import("highlight.js").HLJSApi} hljs 
 * @returns {import("highlight.js").Language}
 */
export function genericlog(hljs) {
    const regex = hljs.regex;

    const TRACE = [
        'TRACE', 'TRAC', 'Trace'
    ]

    const DEBUG = [
        'DEBUG', 'DEBU', 'Debug'
    ]

    const INFO = [
        'INFO', 'Info'
    ]

    const WARN = [
        'WARN', 'Warn', 'WARNING', 'Warning'
    ]
    
    const ERROR = [
        'ERROR', 'ERRO', 'Error', 'Failed'
    ]

    const CRITICAL = [
        'CRITICAL', 'CRIT', 'Critical'
    ]

    return {
        name: 'Generic Log',
        contains: [
            {
                className: 'ip',
                begin: /^\d{1,3}\.\d{1,3}\.\d{1,3}\.\d{1,3}(:\d{1,5})?\b/,
                relevance: 5
            },
            {
                className: 'trace',
                begin: regex.either(...TRACE)+regex.lookahead("(:| )"),
                keywords: TRACE,
                relevance: 1,
            },
            {
                className: 'debug',
                begin: regex.either(...DEBUG)+regex.lookahead("(:| )"),
                relevance: 2,
            },
            {
                className: 'info',
                match: regex.either(...INFO)+regex.lookahead("(:| )"),
                relevance: 3,
            },
            {
                className: 'warn',
                begin: regex.either(...WARN)+regex.lookahead("(:| )"),
                relevance: 4,
            },
            {
                className: 'error',
                begin: regex.either(...ERROR)+regex.lookahead("(:| )"),
                relevance: 5,
            },
            {
                className: 'critical',
                begin: regex.either(...CRITICAL)+regex.lookahead("(:| )"),
                relevance: 6,
            },
        ]
    };
}