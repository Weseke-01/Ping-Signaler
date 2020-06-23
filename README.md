# Ping-Signaler

Mijn ping-signaler programma. Dit programma scant op pings. Dat doet het door het commando 'netstat -s' uit te voeren. Vervolgens gaat er een scanner door de output heen
scannen tot de scanner op 'Messages' stuitert. Er zijn meerdere regels die 'Messages' bevatten die regels worden verzameld. Dan filtert het programma alle andere regels buiten
'Messages' onder ICMPv4 eruit. De waarde die achter deze 'Messages' staat wordt eruit gehaald en vervolgens geprint. Dit proces gebeurd iedere drie seconden.
