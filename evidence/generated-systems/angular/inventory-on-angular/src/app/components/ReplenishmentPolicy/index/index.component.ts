
import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { ReplenishmentPolicyService } from '../../../services/ReplenishmentPolicy.service';
import { ReplenishmentPolicy } from '../../../models/ReplenishmentPolicy';

@Component({
    selector: 'app-index-replenishmentPolicy',
    standalone: false,
    templateUrl: './index.component.html',
    styleUrls: ['./index.component.css']
})
export class IndexReplenishmentPolicyComponent implements OnInit {

    replenishmentPolicys: ReplenishmentPolicy[] = [];

    constructor(
        private router: Router,
        private service: ReplenishmentPolicyService
) {}

    ngOnInit(): void {
        this.getReplenishmentPolicys();
}

    getReplenishmentPolicys(): void {
        this.service.getReplenishmentPolicys().subscribe((res) => {
        this.replenishmentPolicys = res;
    });
}

    deleteReplenishmentPolicy(id: any): void {
        this.service.deleteReplenishmentPolicy(id)
            .subscribe(() => {
                this.getReplenishmentPolicys();
            });
    }
}