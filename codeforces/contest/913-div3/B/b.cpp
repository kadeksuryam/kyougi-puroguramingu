#include<bits/stdc++.h>
using namespace std;

int main() {
    ios_base::sync_with_stdio(0);
    int t; cin >> t;

    stack<pair<char, int>> low, upp;

    while(t--) {
        string s; cin >> s;
        for(int i=0;i<s.size();i++) {
            if(s[i] == 'b') {
                if(!low.empty())
                    low.pop();
                continue;
            } else if(s[i] == 'B') {
               if(!upp.empty())
                    upp.pop();
               continue;
            }

            if(s[i] >= 'a' && s[i] <= 'z') {
                low.push({s[i], i});
            } else {
                upp.push({s[i], i});
            }
        }
        vector<char> after((int)s.size(), '.');
        while(!low.empty()) {
            pair<char, int> curr = low.top();

            after[curr.second] = curr.first;
            low.pop();
        }
        while(!upp.empty()) {
            pair<char, int> curr = upp.top();

            after[curr.second] = curr.first;
            upp.pop();
        }

        string ans = "";

        for(int i=0;i<s.size();i++) {
            if(after[i] != '.') {
                ans += after[i];
            }
        }
        cout << ans << "\n";
    }
    return 0;
}