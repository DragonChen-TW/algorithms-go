Target is compute $Z = (a + bi) * (c + di)$ in 3 real number multiplications.

$$
Z = (a + bi) * (c + di) \\
 = ac - bd + (ad + bc) i
$$

Define

$$
P_1 = (a + b) * (c + d) \\
= ac + ad + bc + bd \\
P_2 = ac \\
P_3 = bd \\
\\
Z = ac - bd + (ad + bc) i \\
= (P_2 - P_3) + (P_1 - P_2 - P_3) i
$$

3 multiplications and 5 additions are needed.