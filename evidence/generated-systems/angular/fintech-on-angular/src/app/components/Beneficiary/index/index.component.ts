
import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { BeneficiaryService } from '../../../services/Beneficiary.service';
import { Beneficiary } from '../../../models/Beneficiary';

@Component({
    selector: 'app-index-beneficiary',
    standalone: false,
    templateUrl: './index.component.html',
    styleUrls: ['./index.component.css']
})
export class IndexBeneficiaryComponent implements OnInit {

    beneficiarys: Beneficiary[] = [];

    constructor(
        private router: Router,
        private service: BeneficiaryService
) {}

    ngOnInit(): void {
        this.getBeneficiarys();
}

    getBeneficiarys(): void {
        this.service.getBeneficiarys().subscribe((res) => {
        this.beneficiarys = res;
    });
}

    deleteBeneficiary(id: any): void {
        this.service.deleteBeneficiary(id)
            .subscribe(() => {
                this.getBeneficiarys();
            });
    }
}