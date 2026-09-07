
import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { InvestmentAccountService } from '../../../services/InvestmentAccount.service';
import { InvestmentAccount } from '../../../models/InvestmentAccount';

@Component({
    selector: 'app-index-investmentAccount',
    standalone: false,
    templateUrl: './index.component.html',
    styleUrls: ['./index.component.css']
})
export class IndexInvestmentAccountComponent implements OnInit {

    investmentAccounts: InvestmentAccount[] = [];

    constructor(
        private router: Router,
        private service: InvestmentAccountService
) {}

    ngOnInit(): void {
        this.getInvestmentAccounts();
}

    getInvestmentAccounts(): void {
        this.service.getInvestmentAccounts().subscribe((res) => {
        this.investmentAccounts = res;
    });
}

    deleteInvestmentAccount(id: any): void {
        this.service.deleteInvestmentAccount(id)
            .subscribe(() => {
                this.getInvestmentAccounts();
            });
    }
}