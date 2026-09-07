
import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { PolicyService } from '../../../services/Policy.service';
import { Policy } from '../../../models/Policy';

@Component({
    selector: 'app-index-policy',
    standalone: false,
    templateUrl: './index.component.html',
    styleUrls: ['./index.component.css']
})
export class IndexPolicyComponent implements OnInit {

    policys: Policy[] = [];

    constructor(
        private router: Router,
        private service: PolicyService
) {}

    ngOnInit(): void {
        this.getPolicys();
}

    getPolicys(): void {
        this.service.getPolicys().subscribe((res) => {
        this.policys = res;
    });
}

    deletePolicy(id: any): void {
        this.service.deletePolicy(id)
            .subscribe(() => {
                this.getPolicys();
            });
    }
}