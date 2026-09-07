
import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { UsageLimitService } from '../../../services/UsageLimit.service';
import { UsageLimit } from '../../../models/UsageLimit';

@Component({
    selector: 'app-index-usageLimit',
    standalone: false,
    templateUrl: './index.component.html',
    styleUrls: ['./index.component.css']
})
export class IndexUsageLimitComponent implements OnInit {

    usageLimits: UsageLimit[] = [];

    constructor(
        private router: Router,
        private service: UsageLimitService
) {}

    ngOnInit(): void {
        this.getUsageLimits();
}

    getUsageLimits(): void {
        this.service.getUsageLimits().subscribe((res) => {
        this.usageLimits = res;
    });
}

    deleteUsageLimit(id: any): void {
        this.service.deleteUsageLimit(id)
            .subscribe(() => {
                this.getUsageLimits();
            });
    }
}