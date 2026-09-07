
import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { ExperimentVariantService } from '../../../services/ExperimentVariant.service';
import { ExperimentVariant } from '../../../models/ExperimentVariant';

@Component({
    selector: 'app-index-experimentVariant',
    standalone: false,
    templateUrl: './index.component.html',
    styleUrls: ['./index.component.css']
})
export class IndexExperimentVariantComponent implements OnInit {

    experimentVariants: ExperimentVariant[] = [];

    constructor(
        private router: Router,
        private service: ExperimentVariantService
) {}

    ngOnInit(): void {
        this.getExperimentVariants();
}

    getExperimentVariants(): void {
        this.service.getExperimentVariants().subscribe((res) => {
        this.experimentVariants = res;
    });
}

    deleteExperimentVariant(id: any): void {
        this.service.deleteExperimentVariant(id)
            .subscribe(() => {
                this.getExperimentVariants();
            });
    }
}