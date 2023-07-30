# -*- coding: utf-8 -*-
"""
Created on Fri Jul 28 16:39:40 2023

@author: prsnl
"""
UserInput = input("Please provide an integer: ")

convertNum = int(UserInput)

if converNum > 0:
    

for i in range(1, convertNum+1):
    if i %3 == 0 and i%5 == 0:
        print("BANG")
    elif i%3 == 0:
        print("weasel")
    elif i%5 == 0:
        print("pop")
    else:
        print(i)
        

        
