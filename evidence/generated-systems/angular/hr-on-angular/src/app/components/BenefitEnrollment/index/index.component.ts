
import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { BenefitEnrollmentService } from '../../../services/BenefitEnrollment.service';
import { BenefitEnrollment } from '../../../models/BenefitEnrollment';

@Component({
    selector: 'app-index-benefitEnrollment',
    standalone: false,
    templateUrl: './index.component.html',
    styleUrls: ['./index.component.css']
})
export class IndexBenefitEnrollmentComponent implements OnInit {

    benefitEnrollments: BenefitEnrollment[] = [];

    constructor(
        private router: Router,
        private service: BenefitEnrollmentService
) {}

    ngOnInit(): void {
        this.getBenefitEnrollments();
}

    getBenefitEnrollments(): void {
        this.service.getBenefitEnrollments().subscribe((res) => {
        this.benefitEnrollments = res;
    });
}

    deleteBenefitEnrollment(id: any): void {
        this.service.deleteBenefitEnrollment(id)
            .subscribe(() => {
                this.getBenefitEnrollments();
            });
    }
}