
import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { AppliedFeeService } from '../../../services/AppliedFee.service';
import { AppliedFee } from '../../../models/AppliedFee';

@Component({
    selector: 'app-index-appliedFee',
    standalone: false,
    templateUrl: './index.component.html',
    styleUrls: ['./index.component.css']
})
export class IndexAppliedFeeComponent implements OnInit {

    appliedFees: AppliedFee[] = [];

    constructor(
        private router: Router,
        private service: AppliedFeeService
) {}

    ngOnInit(): void {
        this.getAppliedFees();
}

    getAppliedFees(): void {
        this.service.getAppliedFees().subscribe((res) => {
        this.appliedFees = res;
    });
}

    deleteAppliedFee(id: any): void {
        this.service.deleteAppliedFee(id)
            .subscribe(() => {
                this.getAppliedFees();
            });
    }
}