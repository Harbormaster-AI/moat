
import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { ThirdPartyAssessmentService } from '../../../services/ThirdPartyAssessment.service';
import { ThirdPartyAssessment } from '../../../models/ThirdPartyAssessment';

@Component({
    selector: 'app-index-thirdPartyAssessment',
    standalone: false,
    templateUrl: './index.component.html',
    styleUrls: ['./index.component.css']
})
export class IndexThirdPartyAssessmentComponent implements OnInit {

    thirdPartyAssessments: ThirdPartyAssessment[] = [];

    constructor(
        private router: Router,
        private service: ThirdPartyAssessmentService
) {}

    ngOnInit(): void {
        this.getThirdPartyAssessments();
}

    getThirdPartyAssessments(): void {
        this.service.getThirdPartyAssessments().subscribe((res) => {
        this.thirdPartyAssessments = res;
    });
}

    deleteThirdPartyAssessment(id: any): void {
        this.service.deleteThirdPartyAssessment(id)
            .subscribe(() => {
                this.getThirdPartyAssessments();
            });
    }
}