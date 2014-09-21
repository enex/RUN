# RUN

Run ist eine kleine rein funktionale Programmiersrache, welche einfach zu lernen ist. 

### Verarbeitung
    if a == b ($a $b = a) ($a $b = b)
wird übersetzt zu

    ["if a == b", {pattern:"$a $b", body:"a"}, {pattern: "$a $b", body:"b"}]

Darauf werden dann die vorher definierten patterns angewendet.
Ein Pattern wird volgendermaßen definiert
    
    [pattern] = [body]

im foltenden Programm wird dann das pattern durch body ersetzt

## Syntax
Die Komplette Programmiersprache baut auf dem Prinzip pattern matching auf. Das heißt die einzige überhaupt mögliche Funktion in RUN ist das matchen von Pattern und ersetzen durch das Ergebnis. Hierbei ist allerdings die Reihenvolge des pattern matchings zu beachten. Ein Match wird stets ausgehend vom aktuellen scope aufwerts gesucht. Das Pattern matching verhält sich damit ähnlich wie Maccros in anderen Programmiersprachen, stellt allerdings trotz des sehr einfachen syntax viele Funktionen bereit

Die Maximale anzahl an Zeilen is 65535 und die Maximale Zeilenlänge beträgt
255 Zeichen. Falls jemand versuchen sollt eine noch größere Zeile zu konstruieren, wird ein Fehler ausgegeben.

### Scopes

Scopes sind Bereiche in denen die Definitionen die darin gemacht wurden gelten.
Weiter eingerückte bereiche erben von den weniger eingerückten. Bei einer definition die mit einrückung definiert wierd, und daher einen eigenen Scope besitzt ist die Letzte definition oder Anweißung das Ergebnis.

### Definition
Eien definition weißt einem pattern einen wert zu.

#### Patterns

    $[name]...
bedeutet das das definierte pattern öffter gematcht wird

    $[name]:[Bedingung]
ermöglicht das genauere speizifizieren des zu suchenden. Die Bdingungsfunktion wird mit dem Wert aufgerufen und gibt entweder true für match oder fals für kein match aus.

Dieße Patterns können beliebig kombiniert werden und die Bedingungen sind wieder durch patterns definierte Funktionen. Wenn Konstante werte in den Patterns definiert sind werden dieße auch übernommen

#### Body
In diesem teil befindet sich das ergebnis was anstelle des patterns geschreiben wird. Alle symbole, die im Pattern definiert werden werden dort durch den zugehörigen Wert ersetzt.


## Anderes
Es ist außerdem möglich eine art typ-system ein zu führen.


Der Einfachheit halber implementiere ich erst mal noch kein typ-system, es ist aber möglich das späther zusätzliche Grammatik für die Patterns eingeführt wird, die Interface artige zugriffe erlaubt. Programme die vor dießer Änderung geschreiben wurden bleiben dadurch weiterhin funktionsfähig.

Ich bin mir zur Zeit noch nicht sicher wie die Implementierung von Patterns genau sein soll.

Problematisch ist vor allem die Verschachtelung von Pattern, die ich gerne so gut es geht vermeiden Möchte.

Ansatz dazu:
    Eine Definition die wiederum ein Pattern definiert, gibt dießes pattern als Ergebnis zurück, volglich wird die funktion bei Aufruf mit dießem Pattern dahinter aufgerufen.
    Probleme
        Es ist nicht differenzierbar ob ein Pattern als ergebnis-wert zurück gegeben wird, oder nicht.
        Lößung:


## Nebenwirkung

Nebe