
import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { MeasureService } from '../../../services/Measure.service';
import { Measure } from '../../../models/Measure';

@Component({
    selector: 'app-index-measure',
    standalone: false,
    templateUrl: './index.component.html',
    styleUrls: ['./index.component.css']
})
export class IndexMeasureComponent implements OnInit {

    measures: Measure[] = [];

    constructor(
        private router: Router,
        private service: MeasureService
) {}

    ngOnInit(): void {
        this.getMeasures();
}

    getMeasures(): void {
        this.service.getMeasures().subscribe((res) => {
        this.measures = res;
    });
}

    deleteMeasure(id: any): void {
        this.service.deleteMeasure(id)
            .subscribe(() => {
                this.getMeasures();
            });
    }
}