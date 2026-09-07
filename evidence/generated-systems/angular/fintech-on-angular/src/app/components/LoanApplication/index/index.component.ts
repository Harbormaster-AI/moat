
import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { LoanApplicationService } from '../../../services/LoanApplication.service';
import { LoanApplication } from '../../../models/LoanApplication';

@Component({
    selector: 'app-index-loanApplication',
    standalone: false,
    templateUrl: './index.component.html',
    styleUrls: ['./index.component.css']
})
export class IndexLoanApplicationComponent implements OnInit {

    loanApplications: LoanApplication[] = [];

    constructor(
        private router: Router,
        private service: LoanApplicationService
) {}

    ngOnInit(): void {
        this.getLoanApplications();
}

    getLoanApplications(): void {
        this.service.getLoanApplications().subscribe((res) => {
        this.loanApplications = res;
    });
}

    deleteLoanApplication(id: any): void {
        this.service.deleteLoanApplication(id)
            .subscribe(() => {
                this.getLoanApplications();
            });
    }
}