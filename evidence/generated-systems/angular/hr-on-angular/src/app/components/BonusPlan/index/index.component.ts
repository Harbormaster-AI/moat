
import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { BonusPlanService } from '../../../services/BonusPlan.service';
import { BonusPlan } from '../../../models/BonusPlan';

@Component({
    selector: 'app-index-bonusPlan',
    standalone: false,
    templateUrl: './index.component.html',
    styleUrls: ['./index.component.css']
})
export class IndexBonusPlanComponent implements OnInit {

    bonusPlans: BonusPlan[] = [];

    constructor(
        private router: Router,
        private service: BonusPlanService
) {}

    ngOnInit(): void {
        this.getBonusPlans();
}

    getBonusPlans(): void {
        this.service.getBonusPlans().subscribe((res) => {
        this.bonusPlans = res;
    });
}

    deleteBonusPlan(id: any): void {
        this.service.deleteBonusPlan(id)
            .subscribe(() => {
                this.getBonusPlans();
            });
    }
}