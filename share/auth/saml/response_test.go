package saml

import (
	"testing"
)

func TestParseResponse(t *testing.T) {
	query := "SAMLResponse=PD94bWwgdmVyc2lvbj0iMS4wIiBlbmNvZGluZz0iVVRGLTgiPz48c2FtbDJwOlJlc3BvbnNlIHhtbG5zOnNhbWwycD0idXJuOm9hc2lzOm5hbWVzOnRjOlNBTUw6Mi4wOnByb3RvY29sIiBEZXN0aW5hdGlvbj0iaHR0cHM6Ly8xMC4yMTEuNTUuMTA6ODQ0My90b2tlbl9hdXRoX3NlcnZlciIgSUQ9ImlkMTE1ODQ5NzQxODI5NzQ2NDgxMTYxNzUyMzUxIiBJblJlc3BvbnNlVG89Il9jNDc4NmJmYi03ZWVkLTQ4NmMtNDZhNS05NWI0YmExNTQzZmIiIElzc3VlSW5zdGFudD0iMjAxNy0wOS0yMFQyMjo0NzowNi4wMjRaIiBWZXJzaW9uPSIyLjAiIHhtbG5zOnhzPSJodHRwOi8vd3d3LnczLm9yZy8yMDAxL1hNTFNjaGVtYSI%2BPHNhbWwyOklzc3VlciB4bWxuczpzYW1sMj0idXJuOm9hc2lzOm5hbWVzOnRjOlNBTUw6Mi4wOmFzc2VydGlvbiIgRm9ybWF0PSJ1cm46b2FzaXM6bmFtZXM6dGM6U0FNTDoyLjA6bmFtZWlkLWZvcm1hdDplbnRpdHkiPmh0dHA6Ly93d3cub2t0YS5jb20vZXhrYmpncDlhMDRZS0l2cW8waDc8L3NhbWwyOklzc3Vlcj48ZHM6U2lnbmF0dXJlIHhtbG5zOmRzPSJodHRwOi8vd3d3LnczLm9yZy8yMDAwLzA5L3htbGRzaWcjIj48ZHM6U2lnbmVkSW5mbz48ZHM6Q2Fub25pY2FsaXphdGlvbk1ldGhvZCBBbGdvcml0aG09Imh0dHA6Ly93d3cudzMub3JnLzIwMDEvMTAveG1sLWV4Yy1jMTRuIyIvPjxkczpTaWduYXR1cmVNZXRob2QgQWxnb3JpdGhtPSJodHRwOi8vd3d3LnczLm9yZy8yMDAxLzA0L3htbGRzaWctbW9yZSNyc2Etc2hhMjU2Ii8%2BPGRzOlJlZmVyZW5jZSBVUkk9IiNpZDExNTg0OTc0MTgyOTc0NjQ4MTE2MTc1MjM1MSI%2BPGRzOlRyYW5zZm9ybXM%2BPGRzOlRyYW5zZm9ybSBBbGdvcml0aG09Imh0dHA6Ly93d3cudzMub3JnLzIwMDAvMDkveG1sZHNpZyNlbnZlbG9wZWQtc2lnbmF0dXJlIi8%2BPGRzOlRyYW5zZm9ybSBBbGdvcml0aG09Imh0dHA6Ly93d3cudzMub3JnLzIwMDEvMTAveG1sLWV4Yy1jMTRuIyI%2BPGVjOkluY2x1c2l2ZU5hbWVzcGFjZXMgeG1sbnM6ZWM9Imh0dHA6Ly93d3cudzMub3JnLzIwMDEvMTAveG1sLWV4Yy1jMTRuIyIgUHJlZml4TGlzdD0ieHMiLz48L2RzOlRyYW5zZm9ybT48L2RzOlRyYW5zZm9ybXM%2BPGRzOkRpZ2VzdE1ldGhvZCBBbGdvcml0aG09Imh0dHA6Ly93d3cudzMub3JnLzIwMDEvMDQveG1sZW5jI3NoYTI1NiIvPjxkczpEaWdlc3RWYWx1ZT5TRkNIamVhc3ZXREJDTXFxcHA3clkwOGVXbDhpdDJib0lrQTBHMUtZU0hNPTwvZHM6RGlnZXN0VmFsdWU%2BPC9kczpSZWZlcmVuY2U%2BPC9kczpTaWduZWRJbmZvPjxkczpTaWduYXR1cmVWYWx1ZT5NNU5lMDJyNDBKdjZlOHNYWmZWd0Z6RjI4empRSERXaXp4Y0lmTlRBVHRZV2tZNlI3aHlxbnJoY1BRbHpnNUJPcDAvdldOSTIwdXdWZjhYZU03Vkd0Y1N1bHBmZE5pYjRMV00vOFhFOGV5NFBLVmVIanBmbWNBdVkvK2V5bHRBaE1GTEpkZEIxTzZmQjFwQm9GTTZXLzVYcFlxOXE3OGdGNitNZVNDTGlJUU40d3RHK3pSSlVrcjJsR0NDU05mUWtLdnFvZXZqUGRIcHZCSmhFQVhpcGRRRUVRb29LK29kd1dXdDVtY1EyMDQxcWx2eGVIbGNGYjZOQk1uUkIwSDU5SWpoMmtZeEZHcitramsyMkxkd3RFM1RuVGVYTFJ1Q1k4VVk5MEpOYlVnVE94WEJhYTJhMTR2NzBjM25UZ3JJZmIrUHloYUlZVlN6STcwRE84Vkora2c9PTwvZHM6U2lnbmF0dXJlVmFsdWU%2BPGRzOktleUluZm8%2BPGRzOlg1MDlEYXRhPjxkczpYNTA5Q2VydGlmaWNhdGU%2BTUlJRHBEQ0NBb3lnQXdJQkFnSUdBVjIrL3BEK01BMEdDU3FHU0liM0RRRUJDd1VBTUlHU01Rc3dDUVlEVlFRR0V3SlZVekVUTUJFRwpBMVVFQ0F3S1EyRnNhV1p2Y201cFlURVdNQlFHQTFVRUJ3d05VMkZ1SUVaeVlXNWphWE5qYnpFTk1Bc0dBMVVFQ2d3RVQydDBZVEVVCk1CSUdBMVVFQ3d3TFUxTlBVSEp2ZG1sa1pYSXhFekFSQmdOVkJBTU1DbVJsZGkweU5UVXpPRGd4SERBYUJna3Foa2lHOXcwQkNRRVcKRFdsdVptOUFiMnQwWVM1amIyMHdIaGNOTVRjd09EQTNNak14T0RBd1doY05NamN3T0RBM01qTXhPVEF3V2pDQmtqRUxNQWtHQTFVRQpCaE1DVlZNeEV6QVJCZ05WQkFnTUNrTmhiR2xtYjNKdWFXRXhGakFVQmdOVkJBY01EVk5oYmlCR2NtRnVZMmx6WTI4eERUQUxCZ05WCkJBb01CRTlyZEdFeEZEQVNCZ05WQkFzTUMxTlRUMUJ5YjNacFpHVnlNUk13RVFZRFZRUUREQXBrWlhZdE1qVTFNemc0TVJ3d0dnWUoKS29aSWh2Y05BUWtCRmcxcGJtWnZRRzlyZEdFdVkyOXRNSUlCSWpBTkJna3Foa2lHOXcwQkFRRUZBQU9DQVE4QU1JSUJDZ0tDQVFFQQppMERadVhsMGZkYkJpdFZOS3ZUajVEc3FkTXFkT3RhSGh4ZWxoSUxRVCtkek1iVzkwOGJUVW5NaEVDaFdtWFZxTjJKakZ6ZXkycEs1CmJzY2U2T0tHWTVYZSthc0c0V01ka2tIeWE2SjA0SW1nRlU0a1Q2cVlsNCt2cTN5R01zT2FHOFQ3TnNtY00wR2xMWklqYkVmT0N1Z1AKa043MTRzeURwemNKY0VmSmszV3E5RlNEVE1WRUNRaGpLM3NaRy9yNGlQNThTWmxUaEhBL1hYTWZGMHl5QzRRWEVGNnMrRFBzdGVGegpjeDIrQjh2aVNVR1p2R010eXNKMi9sTjFuazlxNU9RRWE0MkptUDBDdHpmQTU1UWJMWmtYdFNlR0p5d3FOMTBheWdGS0JBN3ZlMHF6Cmk2UTM5Uk5ndG84MTNvbFU3L3ozVlRsckFzZGhaWHJpMHhMUVFRSURBUUFCTUEwR0NTcUdTSWIzRFFFQkN3VUFBNElCQVFCcjJNbnAKR1ZmVjdkTVdobWJ0ZHJPTms0VW1VKysrVHRMT2cvS3Fxc1BxVlc5Zy9HM2xZSHFubmpTTGR4OEVLSk5KTUFOK3VKRllERGZjSHowYgpKU3lTRngrTGQrUkUwRjJLOURHTEliNDc2ZmhScm1LZTE0aXZ6Qm50ZGtvUHhrQnlXU0NScm9IMlAwQmh6U3dIZUlGTTNlQmR2RktFCnZQYXFON2owWFF2RnR5dGQwM0hsRU5jNEtBZExLbENZcjBIYWs0RVlRMXRadUJrK3p3RmVIUE0rYkZ5KzJlZjRDdmgyZEJsbzFhMFcKZStOaDNvUTlEdjlaWFJXZzVPTUxYL053ZnYyaVdMZW5EWmUzWkFoVGVzRWpyUXJxbThGT3REajBLbnJOVVYwZGErYVRHOTNaTXF4VQpta2RaNHJxeEhEaW1WZHNCQ2VWb3kwMXdlbUxLL2FGaDwvZHM6WDUwOUNlcnRpZmljYXRlPjwvZHM6WDUwOURhdGE%2BPC9kczpLZXlJbmZvPjwvZHM6U2lnbmF0dXJlPjxzYW1sMnA6U3RhdHVzIHhtbG5zOnNhbWwycD0idXJuOm9hc2lzOm5hbWVzOnRjOlNBTUw6Mi4wOnByb3RvY29sIj48c2FtbDJwOlN0YXR1c0NvZGUgVmFsdWU9InVybjpvYXNpczpuYW1lczp0YzpTQU1MOjIuMDpzdGF0dXM6U3VjY2VzcyIvPjwvc2FtbDJwOlN0YXR1cz48c2FtbDI6QXNzZXJ0aW9uIHhtbG5zOnNhbWwyPSJ1cm46b2FzaXM6bmFtZXM6dGM6U0FNTDoyLjA6YXNzZXJ0aW9uIiBJRD0iaWQxMTU4NDk3NDE4MzAwMTgzOTQ0NTAwODQ1NCIgSXNzdWVJbnN0YW50PSIyMDE3LTA5LTIwVDIyOjQ3OjA2LjAyNFoiIFZlcnNpb249IjIuMCIgeG1sbnM6eHM9Imh0dHA6Ly93d3cudzMub3JnLzIwMDEvWE1MU2NoZW1hIj48c2FtbDI6SXNzdWVyIEZvcm1hdD0idXJuOm9hc2lzOm5hbWVzOnRjOlNBTUw6Mi4wOm5hbWVpZC1mb3JtYXQ6ZW50aXR5IiB4bWxuczpzYW1sMj0idXJuOm9hc2lzOm5hbWVzOnRjOlNBTUw6Mi4wOmFzc2VydGlvbiI%2BaHR0cDovL3d3dy5va3RhLmNvbS9leGtiamdwOWEwNFlLSXZxbzBoNzwvc2FtbDI6SXNzdWVyPjxkczpTaWduYXR1cmUgeG1sbnM6ZHM9Imh0dHA6Ly93d3cudzMub3JnLzIwMDAvMDkveG1sZHNpZyMiPjxkczpTaWduZWRJbmZvPjxkczpDYW5vbmljYWxpemF0aW9uTWV0aG9kIEFsZ29yaXRobT0iaHR0cDovL3d3dy53My5vcmcvMjAwMS8xMC94bWwtZXhjLWMxNG4jIi8%2BPGRzOlNpZ25hdHVyZU1ldGhvZCBBbGdvcml0aG09Imh0dHA6Ly93d3cudzMub3JnLzIwMDEvMDQveG1sZHNpZy1tb3JlI3JzYS1zaGEyNTYiLz48ZHM6UmVmZXJlbmNlIFVSST0iI2lkMTE1ODQ5NzQxODMwMDE4Mzk0NDUwMDg0NTQiPjxkczpUcmFuc2Zvcm1zPjxkczpUcmFuc2Zvcm0gQWxnb3JpdGhtPSJodHRwOi8vd3d3LnczLm9yZy8yMDAwLzA5L3htbGRzaWcjZW52ZWxvcGVkLXNpZ25hdHVyZSIvPjxkczpUcmFuc2Zvcm0gQWxnb3JpdGhtPSJodHRwOi8vd3d3LnczLm9yZy8yMDAxLzEwL3htbC1leGMtYzE0biMiPjxlYzpJbmNsdXNpdmVOYW1lc3BhY2VzIHhtbG5zOmVjPSJodHRwOi8vd3d3LnczLm9yZy8yMDAxLzEwL3htbC1leGMtYzE0biMiIFByZWZpeExpc3Q9InhzIi8%2BPC9kczpUcmFuc2Zvcm0%2BPC9kczpUcmFuc2Zvcm1zPjxkczpEaWdlc3RNZXRob2QgQWxnb3JpdGhtPSJodHRwOi8vd3d3LnczLm9yZy8yMDAxLzA0L3htbGVuYyNzaGEyNTYiLz48ZHM6RGlnZXN0VmFsdWU%2BOVRCTW5aZlNZdjcrVTBVbnlJNEdwdG8vaFZGZis3bTh2UXNTN1pvVHM1az08L2RzOkRpZ2VzdFZhbHVlPjwvZHM6UmVmZXJlbmNlPjwvZHM6U2lnbmVkSW5mbz48ZHM6U2lnbmF0dXJlVmFsdWU%2BZEMreEZ5RCtBaDlmS3YrSkpjbnZ4MHZsTEUzYXhZZEhkaXd5Wk5aS1NhdVc0UDViTnlBMEkxU1VlVFFOLy9tV0t1OExkb3lPZDl2NFFWblpkMENtTGRFZXRLZUwwQTd3VlZva25Ua00xdWkxbnN3OExScU9Ld0M5ZFpPcWM1Y05PNFpoS1FZbDZZUjY3WmZjVWVXUWl6N0t2VTRJNytsRVFMcENHOXRya1J1dkZUTGhLWFlMSWpVWWVSWmhKcy9pYzMzYS9NKzM2anVsQjFmOTU3eHoxTFVZRW8raXVHRklmay9Ld0tSK1cwWW1rNi9kSUZHQlVpeld3Tm9iN1lsVkp1blR3MjZ6ZTNlU293dWp3dFlGd2V4amlTc05lQTRQaktVdkVISVFFdkVOclF6czk4UU9UbUtHS2JNL2R4NDI3WWRMTmZWYzhKQWM5Z2FSek1iUXFnPT08L2RzOlNpZ25hdHVyZVZhbHVlPjxkczpLZXlJbmZvPjxkczpYNTA5RGF0YT48ZHM6WDUwOUNlcnRpZmljYXRlPk1JSURwRENDQW95Z0F3SUJBZ0lHQVYyKy9wRCtNQTBHQ1NxR1NJYjNEUUVCQ3dVQU1JR1NNUXN3Q1FZRFZRUUdFd0pWVXpFVE1CRUcKQTFVRUNBd0tRMkZzYVdadmNtNXBZVEVXTUJRR0ExVUVCd3dOVTJGdUlFWnlZVzVqYVhOamJ6RU5NQXNHQTFVRUNnd0VUMnQwWVRFVQpNQklHQTFVRUN3d0xVMU5QVUhKdmRtbGtaWEl4RXpBUkJnTlZCQU1NQ21SbGRpMHlOVFV6T0RneEhEQWFCZ2txaGtpRzl3MEJDUUVXCkRXbHVabTlBYjJ0MFlTNWpiMjB3SGhjTk1UY3dPREEzTWpNeE9EQXdXaGNOTWpjd09EQTNNak14T1RBd1dqQ0JrakVMTUFrR0ExVUUKQmhNQ1ZWTXhFekFSQmdOVkJBZ01Da05oYkdsbWIzSnVhV0V4RmpBVUJnTlZCQWNNRFZOaGJpQkdjbUZ1WTJselkyOHhEVEFMQmdOVgpCQW9NQkU5cmRHRXhGREFTQmdOVkJBc01DMU5UVDFCeWIzWnBaR1Z5TVJNd0VRWURWUVFEREFwa1pYWXRNalUxTXpnNE1Sd3dHZ1lKCktvWklodmNOQVFrQkZnMXBibVp2UUc5cmRHRXVZMjl0TUlJQklqQU5CZ2txaGtpRzl3MEJBUUVGQUFPQ0FROEFNSUlCQ2dLQ0FRRUEKaTBEWnVYbDBmZGJCaXRWTkt2VGo1RHNxZE1xZE90YUhoeGVsaElMUVQrZHpNYlc5MDhiVFVuTWhFQ2hXbVhWcU4ySmpGemV5MnBLNQpic2NlNk9LR1k1WGUrYXNHNFdNZGtrSHlhNkowNEltZ0ZVNGtUNnFZbDQrdnEzeUdNc09hRzhUN05zbWNNMEdsTFpJamJFZk9DdWdQCmtONzE0c3lEcHpjSmNFZkprM1dxOUZTRFRNVkVDUWhqSzNzWkcvcjRpUDU4U1psVGhIQS9YWE1mRjB5eUM0UVhFRjZzK0RQc3RlRnoKY3gyK0I4dmlTVUdadkdNdHlzSjIvbE4xbms5cTVPUUVhNDJKbVAwQ3R6ZkE1NVFiTFprWHRTZUdKeXdxTjEwYXlnRktCQTd2ZTBxegppNlEzOVJOZ3RvODEzb2xVNy96M1ZUbHJBc2RoWlhyaTB4TFFRUUlEQVFBQk1BMEdDU3FHU0liM0RRRUJDd1VBQTRJQkFRQnIyTW5wCkdWZlY3ZE1XaG1idGRyT05rNFVtVSsrK1R0TE9nL0txcXNQcVZXOWcvRzNsWUhxbm5qU0xkeDhFS0pOSk1BTit1SkZZRERmY0h6MGIKSlN5U0Z4K0xkK1JFMEYySzlER0xJYjQ3NmZoUnJtS2UxNGl2ekJudGRrb1B4a0J5V1NDUnJvSDJQMEJoelN3SGVJRk0zZUJkdkZLRQp2UGFxTjdqMFhRdkZ0eXRkMDNIbEVOYzRLQWRMS2xDWXIwSGFrNEVZUTF0WnVCayt6d0ZlSFBNK2JGeSsyZWY0Q3ZoMmRCbG8xYTBXCmUrTmgzb1E5RHY5WlhSV2c1T01MWC9Od2Z2MmlXTGVuRFplM1pBaFRlc0VqclFycW04Rk90RGowS25yTlVWMGRhK2FURzkzWk1xeFUKbWtkWjRycXhIRGltVmRzQkNlVm95MDF3ZW1MSy9hRmg8L2RzOlg1MDlDZXJ0aWZpY2F0ZT48L2RzOlg1MDlEYXRhPjwvZHM6S2V5SW5mbz48L2RzOlNpZ25hdHVyZT48c2FtbDI6U3ViamVjdCB4bWxuczpzYW1sMj0idXJuOm9hc2lzOm5hbWVzOnRjOlNBTUw6Mi4wOmFzc2VydGlvbiI%2BPHNhbWwyOk5hbWVJRCBGb3JtYXQ9InVybjpvYXNpczpuYW1lczp0YzpTQU1MOjEuMTpuYW1laWQtZm9ybWF0OnVuc3BlY2lmaWVkIj5hdHVuZ0BuZXV2ZWN0b3IuY29tPC9zYW1sMjpOYW1lSUQ%2BPHNhbWwyOlN1YmplY3RDb25maXJtYXRpb24gTWV0aG9kPSJ1cm46b2FzaXM6bmFtZXM6dGM6U0FNTDoyLjA6Y206YmVhcmVyIj48c2FtbDI6U3ViamVjdENvbmZpcm1hdGlvbkRhdGEgSW5SZXNwb25zZVRvPSJfYzQ3ODZiZmItN2VlZC00ODZjLTQ2YTUtOTViNGJhMTU0M2ZiIiBOb3RPbk9yQWZ0ZXI9IjIwMTctMDktMjBUMjI6NTI6MDYuMDI0WiIgUmVjaXBpZW50PSJodHRwczovLzEwLjIxMS41NS4xMDo4NDQzL3Rva2VuX2F1dGhfc2VydmVyIi8%2BPC9zYW1sMjpTdWJqZWN0Q29uZmlybWF0aW9uPjwvc2FtbDI6U3ViamVjdD48c2FtbDI6Q29uZGl0aW9ucyBOb3RCZWZvcmU9IjIwMTctMDktMjBUMjI6NDI6MDYuMDI0WiIgTm90T25PckFmdGVyPSIyMDE3LTA5LTIwVDIyOjUyOjA2LjAyNFoiIHhtbG5zOnNhbWwyPSJ1cm46b2FzaXM6bmFtZXM6dGM6U0FNTDoyLjA6YXNzZXJ0aW9uIj48c2FtbDI6QXVkaWVuY2VSZXN0cmljdGlvbj48c2FtbDI6QXVkaWVuY2U%2BaHR0cDovLzEwLjEuNS4xNzo1MDAwL3NhbWwvc3NvL2V4YW1wbGUtb2t0YS1jb208L3NhbWwyOkF1ZGllbmNlPjwvc2FtbDI6QXVkaWVuY2VSZXN0cmljdGlvbj48L3NhbWwyOkNvbmRpdGlvbnM%2BPHNhbWwyOkF1dGhuU3RhdGVtZW50IEF1dGhuSW5zdGFudD0iMjAxNy0wOS0yMFQyMTozOTo1MS4yNjlaIiBTZXNzaW9uSW5kZXg9Il9jNDc4NmJmYi03ZWVkLTQ4NmMtNDZhNS05NWI0YmExNTQzZmIiIHhtbG5zOnNhbWwyPSJ1cm46b2FzaXM6bmFtZXM6dGM6U0FNTDoyLjA6YXNzZXJ0aW9uIj48c2FtbDI6QXV0aG5Db250ZXh0PjxzYW1sMjpBdXRobkNvbnRleHRDbGFzc1JlZj51cm46b2FzaXM6bmFtZXM6dGM6U0FNTDoyLjA6YWM6Y2xhc3NlczpQYXNzd29yZFByb3RlY3RlZFRyYW5zcG9ydDwvc2FtbDI6QXV0aG5Db250ZXh0Q2xhc3NSZWY%2BPC9zYW1sMjpBdXRobkNvbnRleHQ%2BPC9zYW1sMjpBdXRoblN0YXRlbWVudD48c2FtbDI6QXR0cmlidXRlU3RhdGVtZW50IHhtbG5zOnNhbWwyPSJ1cm46b2FzaXM6bmFtZXM6dGM6U0FNTDoyLjA6YXNzZXJ0aW9uIj48c2FtbDI6QXR0cmlidXRlIE5hbWU9IkZpcnN0TmFtZSIgTmFtZUZvcm1hdD0idXJuOm9hc2lzOm5hbWVzOnRjOlNBTUw6Mi4wOmF0dHJuYW1lLWZvcm1hdDp1bnNwZWNpZmllZCI%2BPHNhbWwyOkF0dHJpYnV0ZVZhbHVlIHhtbG5zOnhzPSJodHRwOi8vd3d3LnczLm9yZy8yMDAxL1hNTFNjaGVtYSIgeG1sbnM6eHNpPSJodHRwOi8vd3d3LnczLm9yZy8yMDAxL1hNTFNjaGVtYS1pbnN0YW5jZSIgeHNpOnR5cGU9InhzOnN0cmluZyI%2BYW5kc29uPC9zYW1sMjpBdHRyaWJ1dGVWYWx1ZT48L3NhbWwyOkF0dHJpYnV0ZT48c2FtbDI6QXR0cmlidXRlIE5hbWU9Ikxhc3ROYW1lIiBOYW1lRm9ybWF0PSJ1cm46b2FzaXM6bmFtZXM6dGM6U0FNTDoyLjA6YXR0cm5hbWUtZm9ybWF0OnVuc3BlY2lmaWVkIj48c2FtbDI6QXR0cmlidXRlVmFsdWUgeG1sbnM6eHM9Imh0dHA6Ly93d3cudzMub3JnLzIwMDEvWE1MU2NoZW1hIiB4bWxuczp4c2k9Imh0dHA6Ly93d3cudzMub3JnLzIwMDEvWE1MU2NoZW1hLWluc3RhbmNlIiB4c2k6dHlwZT0ieHM6c3RyaW5nIj50dW5nPC9zYW1sMjpBdHRyaWJ1dGVWYWx1ZT48L3NhbWwyOkF0dHJpYnV0ZT48c2FtbDI6QXR0cmlidXRlIE5hbWU9IkVtYWlsIiBOYW1lRm9ybWF0PSJ1cm46b2FzaXM6bmFtZXM6dGM6U0FNTDoyLjA6YXR0cm5hbWUtZm9ybWF0OnVuc3BlY2lmaWVkIj48c2FtbDI6QXR0cmlidXRlVmFsdWUgeG1sbnM6eHM9Imh0dHA6Ly93d3cudzMub3JnLzIwMDEvWE1MU2NoZW1hIiB4bWxuczp4c2k9Imh0dHA6Ly93d3cudzMub3JnLzIwMDEvWE1MU2NoZW1hLWluc3RhbmNlIiB4c2k6dHlwZT0ieHM6c3RyaW5nIj5hdHVuZ0BuZXV2ZWN0b3IuY29tPC9zYW1sMjpBdHRyaWJ1dGVWYWx1ZT48L3NhbWwyOkF0dHJpYnV0ZT48c2FtbDI6QXR0cmlidXRlIE5hbWU9IkFkbWluR3JvdXAiIE5hbWVGb3JtYXQ9InVybjpvYXNpczpuYW1lczp0YzpTQU1MOjIuMDphdHRybmFtZS1mb3JtYXQ6dW5zcGVjaWZpZWQiPjxzYW1sMjpBdHRyaWJ1dGVWYWx1ZSB4bWxuczp4cz0iaHR0cDovL3d3dy53My5vcmcvMjAwMS9YTUxTY2hlbWEiIHhtbG5zOnhzaT0iaHR0cDovL3d3dy53My5vcmcvMjAwMS9YTUxTY2hlbWEtaW5zdGFuY2UiIHhzaTp0eXBlPSJ4czpzdHJpbmciPmFkbWluPC9zYW1sMjpBdHRyaWJ1dGVWYWx1ZT48L3NhbWwyOkF0dHJpYnV0ZT48c2FtbDI6QXR0cmlidXRlIE5hbWU9IlJlYWRlckdyb3VwIiBOYW1lRm9ybWF0PSJ1cm46b2FzaXM6bmFtZXM6dGM6U0FNTDoyLjA6YXR0cm5hbWUtZm9ybWF0OnVuc3BlY2lmaWVkIj48c2FtbDI6QXR0cmlidXRlVmFsdWUgeG1sbnM6eHM9Imh0dHA6Ly93d3cudzMub3JnLzIwMDEvWE1MU2NoZW1hIiB4bWxuczp4c2k9Imh0dHA6Ly93d3cudzMub3JnLzIwMDEvWE1MU2NoZW1hLWluc3RhbmNlIiB4c2k6dHlwZT0ieHM6c3RyaW5nIj5yZWFkZXI8L3NhbWwyOkF0dHJpYnV0ZVZhbHVlPjwvc2FtbDI6QXR0cmlidXRlPjwvc2FtbDI6QXR0cmlidXRlU3RhdGVtZW50Pjwvc2FtbDI6QXNzZXJ0aW9uPjwvc2FtbDJwOlJlc3BvbnNlPg%3D%3D&RelayState="

	sp := ServiceProvider{
		IDPSSOURL:           "https://dev-255388.oktapreview.com/app/neuvectordev255388_examplesamlapplication_1/exkbjgp9a04YKIvqo0h7/sso/saml",
		IDPSSODescriptorURL: "http://www.okta.com/exkbjgp9a04YKIvqo0h7",
		IDPPublicCert: `-----BEGIN CERTIFICATE-----
MIIDpDCCAoygAwIBAgIGAV2+/pD+MA0GCSqGSIb3DQEBCwUAMIGSMQswCQYDVQQGEwJVUzETMBEG
A1UECAwKQ2FsaWZvcm5pYTEWMBQGA1UEBwwNU2FuIEZyYW5jaXNjbzENMAsGA1UECgwET2t0YTEU
MBIGA1UECwwLU1NPUHJvdmlkZXIxEzARBgNVBAMMCmRldi0yNTUzODgxHDAaBgkqhkiG9w0BCQEW
DWluZm9Ab2t0YS5jb20wHhcNMTcwODA3MjMxODAwWhcNMjcwODA3MjMxOTAwWjCBkjELMAkGA1UE
BhMCVVMxEzARBgNVBAgMCkNhbGlmb3JuaWExFjAUBgNVBAcMDVNhbiBGcmFuY2lzY28xDTALBgNV
BAoMBE9rdGExFDASBgNVBAsMC1NTT1Byb3ZpZGVyMRMwEQYDVQQDDApkZXYtMjU1Mzg4MRwwGgYJ
KoZIhvcNAQkBFg1pbmZvQG9rdGEuY29tMIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEA
i0DZuXl0fdbBitVNKvTj5DsqdMqdOtaHhxelhILQT+dzMbW908bTUnMhEChWmXVqN2JjFzey2pK5
bsce6OKGY5Xe+asG4WMdkkHya6J04ImgFU4kT6qYl4+vq3yGMsOaG8T7NsmcM0GlLZIjbEfOCugP
kN714syDpzcJcEfJk3Wq9FSDTMVECQhjK3sZG/r4iP58SZlThHA/XXMfF0yyC4QXEF6s+DPsteFz
cx2+B8viSUGZvGMtysJ2/lN1nk9q5OQEa42JmP0CtzfA55QbLZkXtSeGJywqN10aygFKBA7ve0qz
i6Q39RNgto813olU7/z3VTlrAsdhZXri0xLQQQIDAQABMA0GCSqGSIb3DQEBCwUAA4IBAQBr2Mnp
GVfV7dMWhmbtdrONk4UmU+++TtLOg/KqqsPqVW9g/G3lYHqnnjSLdx8EKJNJMAN+uJFYDDfcHz0b
JSySFx+Ld+RE0F2K9DGLIb476fhRrmKe14ivzBntdkoPxkByWSCRroH2P0BhzSwHeIFM3eBdvFKE
vPaqN7j0XQvFtytd03HlENc4KAdLKlCYr0Hak4EYQ1tZuBk+zwFeHPM+bFy+2ef4Cvh2dBlo1a0W
e+Nh3oQ9Dv9ZXRWg5OMLX/Nwfv2iWLenDZe3ZAhTesEjrQrqm8FOtDj0KnrNUV0da+aTG93ZMqxU
mkdZ4rqxHDimVdsBCeVoy01wemLK/aFh
-----END CERTIFICATE-----
`,
	}

	resp, err := ParseSAMLResponse(query)
	if err != nil {
		t.Errorf("Parse SAML response error: %v\n", err.Error())
	}

	err = resp.Validate(&sp, false)
	if err != nil {
		t.Errorf("Validate SAML signature error: %v\n", err.Error())
	}

	attrs := resp.GetAttributes()
	if v, ok := attrs["LastName"]; !ok || len(v) != 1 || v[0] != "tung" {
		t.Errorf("Error in attributes: %v\n", attrs)
	}
}
