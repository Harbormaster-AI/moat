
import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { ExperimentService } from '../../../services/Experiment.service';
import { Experiment } from '../../../models/Experiment';

@Component({
    selector: 'app-index-experiment',
    standalone: false,
    templateUrl: './index.component.html',
    styleUrls: ['./index.component.css']
})
export class IndexExperimentComponent implements OnInit {

    experiments: Experiment[] = [];

    constructor(
        private router: Router,
        private service: ExperimentService
) {}

    ngOnInit(): void {
        this.getExperiments();
}

    getExperiments(): void {
        this.service.getExperiments().subscribe((res) => {
        this.experiments = res;
    });
}

    deleteExperiment(id: any): void {
        this.service.deleteExperiment(id)
            .subscribe(() => {
                this.getExperiments();
            });
    }
}