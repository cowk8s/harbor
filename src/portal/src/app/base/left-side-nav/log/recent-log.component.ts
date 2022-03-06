import { Component, OnInit } from '@angular/core';
import { AuditlogService } from "../../../../../ng-swagger-gen/services/auditlog.service";
import { AuditLog } from "../../../../../ng-swagger-gen/models/audit-log";

@Component({
    selector: 'hbr-log',
    templateUrl: './recent-log.component.html',
})
export class RecentLogComponent implements OnInit {
    recentLogs: AuditLog[] = [];
    ngOnInit(): void {
    }

    
}