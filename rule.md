$$
FIRST(N) = \bigcup_{N \rightarrow \alpha}  FIRST(\alpha) \\[5pt]
FIRST(\epsilon) = \{\epsilon\} \\[5pt]
FIRST(x\beta) = \{x\} \\[5pt]
FIRST(N\beta) = FIRST(N) \bullet FIRST(\beta) \\
~\\ % 插入空白行
FOLLOW(N) = \bigcup_{N' \rightarrow \alpha N\beta}  FIRST(\beta) \bullet FOLLOW(N')
\\
~\\ % 插入空白行
PREDICT(N \rightarrow \alpha) = FIRST(\alpha) \bullet FOLLOW(N) \\[10pt]
where \ \ S \bullet S' = (S - {\epsilon}) \cup S' \quad  \epsilon \in S\\
S \bullet S' = S \quad \epsilon \notin S \\
$$