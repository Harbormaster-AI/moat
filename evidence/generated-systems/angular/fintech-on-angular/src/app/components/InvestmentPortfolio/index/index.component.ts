
import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { InvestmentPortfolioService } from '../../../services/InvestmentPortfolio.service';
import { InvestmentPortfolio } from '../../../models/InvestmentPortfolio';

@Component({
    selector: 'app-index-investmentPortfolio',
    standalone: false,
    templateUrl: './index.component.html',
    styleUrls: ['./index.component.css']
})
export class IndexInvestmentPortfolioComponent implements OnInit {

    investmentPortfolios: InvestmentPortfolio[] = [];

    constructor(
        private router: Router,
        private service: InvestmentPortfolioService
) {}

    ngOnInit(): void {
        this.getInvestmentPortfolios();
}

    getInvestmentPortfolios(): void {
        this.service.getInvestmentPortfolios().subscribe((res) => {
        this.investmentPortfolios = res;
    });
}

    deleteInvestmentPortfolio(id: any): void {
        this.service.deleteInvestmentPortfolio(id)
            .subscribe(() => {
                this.getInvestmentPortfolios();
            });
    }
}