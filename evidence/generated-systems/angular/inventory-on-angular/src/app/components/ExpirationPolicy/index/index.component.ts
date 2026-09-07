
import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { ExpirationPolicyService } from '../../../services/ExpirationPolicy.service';
import { ExpirationPolicy } from '../../../models/ExpirationPolicy';

@Component({
    selector: 'app-index-expirationPolicy',
    standalone: false,
    templateUrl: './index.component.html',
    styleUrls: ['./index.component.css']
})
export class IndexExpirationPolicyComponent implements OnInit {

    expirationPolicys: ExpirationPolicy[] = [];

    constructor(
        private router: Router,
        private service: ExpirationPolicyService
) {}

    ngOnInit(): void {
        this.getExpirationPolicys();
}

    getExpirationPolicys(): void {
        this.service.getExpirationPolicys().subscribe((res) => {
        this.expirationPolicys = res;
    });
}

    deleteExpirationPolicy(id: any): void {
        this.service.deleteExpirationPolicy(id)
            .subscribe(() => {
                this.getExpirationPolicys();
            });
    }
}