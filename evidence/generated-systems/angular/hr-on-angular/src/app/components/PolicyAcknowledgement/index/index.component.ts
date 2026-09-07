
import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { PolicyAcknowledgementService } from '../../../services/PolicyAcknowledgement.service';
import { PolicyAcknowledgement } from '../../../models/PolicyAcknowledgement';

@Component({
    selector: 'app-index-policyAcknowledgement',
    standalone: false,
    templateUrl: './index.component.html',
    styleUrls: ['./index.component.css']
})
export class IndexPolicyAcknowledgementComponent implements OnInit {

    policyAcknowledgements: PolicyAcknowledgement[] = [];

    constructor(
        private router: Router,
        private service: PolicyAcknowledgementService
) {}

    ngOnInit(): void {
        this.getPolicyAcknowledgements();
}

    getPolicyAcknowledgements(): void {
        this.service.getPolicyAcknowledgements().subscribe((res) => {
        this.policyAcknowledgements = res;
    });
}

    deletePolicyAcknowledgement(id: any): void {
        this.service.deletePolicyAcknowledgement(id)
            .subscribe(() => {
                this.getPolicyAcknowledgements();
            });
    }
}