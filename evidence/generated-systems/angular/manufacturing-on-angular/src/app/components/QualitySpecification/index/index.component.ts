
import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { QualitySpecificationService } from '../../../services/QualitySpecification.service';
import { QualitySpecification } from '../../../models/QualitySpecification';

@Component({
    selector: 'app-index-qualitySpecification',
    standalone: false,
    templateUrl: './index.component.html',
    styleUrls: ['./index.component.css']
})
export class IndexQualitySpecificationComponent implements OnInit {

    qualitySpecifications: QualitySpecification[] = [];

    constructor(
        private router: Router,
        private service: QualitySpecificationService
) {}

    ngOnInit(): void {
        this.getQualitySpecifications();
}

    getQualitySpecifications(): void {
        this.service.getQualitySpecifications().subscribe((res) => {
        this.qualitySpecifications = res;
    });
}

    deleteQualitySpecification(id: any): void {
        this.service.deleteQualitySpecification(id)
            .subscribe(() => {
                this.getQualitySpecifications();
            });
    }
}