#!/bin/ksh
grep -rl configuration . | xargs sed -i .bckup 's/configuration/configuration/g'
find ./ -name "*.bckup*" -delete
grep -rl remoterestore . | xargs sed -i .bckup 's/remoterestore/remoterestore/g'
find ./ -name "*.bckup*" -delete
grep -rl roles . | xargs sed -i .bckup 's/roles/roles/g'
find ./ -name "*.bckup*" -delete
grep -rl search . | xargs sed -i .bckup 's/search/search/g'
find ./ -name "*.bckup*" -delete
grep -rl smbfileopens . | xargs sed -i .bckup 's/smbfileopens/smbfileopens/g'
find ./ -name "*.bckup*" -delete
grep -rl statistics . | xargs sed -i .bckup 's/statistics/statistics/g'
find ./ -name "*.bckup*" -delete
grep -rl tenants . | xargs sed -i .bckup 's/tenants/tenants/g'
find ./ -name "*.bckup*" -delete
grep -rl vaults . | xargs sed -i .bckup 's/vaults/vaults/g'
find ./ -name "*.bckup*" -delete
grep -rl restoretasks . | xargs sed -i .bckup 's/restoretasks/restoretasks/g'
find ./ -name "*.bckup*" -delete
grep -rl viewboxes . | xargs sed -i .bckup 's/viewboxes/viewboxes/g'
find ./ -name "*.bckup*" -delete
grep -rl views . | xargs sed -i .bckup 's/views/views/g'
find ./ -name "*.bckup*" -delete
grep -rl vlan . | xargs sed -i .bckup 's/vlan/vlan/g'
find ./ -name "*.bckup*" -delete
grep -rl appinstance . | xargs sed -i .bckup 's/appinstance/appinstance/g'
find ./ -name "*.bckup*" -delete
grep -rl app . | xargs sed -i .bckup 's/app/app/g'
find ./ -name "*.bckup*" -delete
grep -rl certificates . | xargs sed -i .bckup 's/certificates/certificates/g'
find ./ -name "*.bckup*" -delete
grep -rl clusters . | xargs sed -i .bckup 's/clusters/clusters/g'
find ./ -name "*.bckup*" -delete
grep -rl interfacegroup . | xargs sed -i .bckup 's/interfacegroup/interfacegroup/g'
find ./ -name "*.bckup*" -delete
grep -rl nodes . | xargs sed -i .bckup 's/nodes/nodes/g'
find ./ -name "*.bckup*" -delete
grep -rl remotecluster . | xargs sed -i .bckup 's/remotecluster/remotecluster/g'
find ./ -name "*.bckup*" -delete
grep -rl routes . | xargs sed -i .bckup 's/routes/routes/g'
find ./ -name "*.bckup*" -delete
grep -rl principals . | xargs sed -i .bckup 's/principals/principals/g'
find ./ -name "*.bckup*" -delete
grep -rl notifications . | xargs sed -i .bckup 's/notifications/notifications/g'
find ./ -name "*.bckup*" -delete
grep -rl preferences . | xargs sed -i .bckup 's/preferences/preferences/g'
find ./ -name "*.bckup*" -delete
grep -rl staticroute . | xargs sed -i .bckup 's/staticroute/staticroute/g'
find ./ -name "*.bckup*" -delete
grep -rl tenant . | xargs sed -i .bckup 's/tenant/tenant/g'
find ./ -name "*.bckup*" -delete
grep -rl activedirectory . | xargs sed -i .bckup 's/activedirectory/activedirectory/g'
find ./ -name "*.bckup*" -delete
grep -rl alerts . | xargs sed -i .bckup 's/alerts/alerts/g'
find ./ -name "*.bckup*" -delete
grep -rl protectionsources . | xargs sed -i .bckup 's/protectionsources/protectionsources/g'
find ./ -name "*.bckup*" -delete
grep -rl protectionruns . | xargs sed -i .bckup 's/protectionruns/protectionruns/g'
find ./ -name "*.bckup*" -delete
grep -rl protectionpolicies . | xargs sed -i .bckup 's/protectionpolicies/protectionpolicies/g'
find ./ -name "*.bckup*" -delete
grep -rl protectionjobs . | xargs sed -i .bckup 's/protectionjobs/protectionjobs/g'
find ./ -name "*.bckup*" -delete
grep -rl audit . | xargs sed -i .bckup 's/audit/audit/g'
find ./ -name "*.bckup*" -delete
grep -rl kmsconfiguration . | xargs sed -i .bckup 's/kmsconfiguration/kmsconfiguration/g'
find ./ -name "*.bckup*" -delete
grep -rl privileges . | xargs sed -i .bckup 's/privileges/privileges/g'
find ./ -name "*.bckup*" -delete
grep -rl ldapprovider . | xargs sed -i .bckup 's/ldapprovider/ldapprovider/g'
find ./ -name "*.bckup*" -delete
grep -rl mimport . | xargs sed -i .bckup 's/mimport/mimport/g'
find ./ -name "*.bckup*" -delete
grep -rl idps . | xargs sed -i .bckup 's/idps/idps/g'
find ./ -name "*.bckup*" -delete
grep -rl groups . | xargs sed -i .bckup 's/groups/groups/g'
find ./ -name "*.bckup*" -delete
grep -rl dashboard . | xargs sed -i .bckup 's/dashboard/dashboard/g'
find ./ -name "*.bckup*" -delete
grep -rl clusterpartitions . | xargs sed -i .bckup 's/clusterpartitions/clusterpartitions/g'
find ./ -name "*.bckup*" -delete
grep -rl export . | xargs sed -i .bckup 's/export/export/g'
find ./ -name "*.bckup*" -delete
grep -rl cluster . | xargs sed -i .bckup 's/cluster/cluster/g'
find ./ -name "*.bckup*" -delete
grep -rl accesstokens . | xargs sed -i .bckup 's/accesstokens/accesstokens/g'
find ./ -name "*.bckup*" -delete
grep -rl models . | xargs sed -i .bckup 's/models/models/g'
find ./ -name "*.bckup*" -delete
