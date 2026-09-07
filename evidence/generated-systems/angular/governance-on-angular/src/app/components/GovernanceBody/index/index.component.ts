
import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { GovernanceBodyService } from '../../../services/GovernanceBody.service';
import { GovernanceBody } from '../../../models/GovernanceBody';

@Component({
    selector: 'app-index-governanceBody',
    standalone: false,
    templateUrl: './index.component.html',
    styleUrls: ['./index.component.css']
})
export class IndexGovernanceBodyComponent implements OnInit {

    governanceBodys: GovernanceBody[] = [];

    constructor(
        private router: Router,
        private service: GovernanceBodyService
) {}

    ngOnInit(): void {
        this.getGovernanceBodys();
}

    getGovernanceBodys(): void {
        this.service.getGovernanceBodys().subscribe((res) => {
        this.governanceBodys = res;
    });
}

    deleteGovernanceBody(id: any): void {
        this.service.deleteGovernanceBody(id)
            .subscribe(() => {
                this.getGovernanceBodys();
            });
    }
}