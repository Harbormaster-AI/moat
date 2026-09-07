
import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { TimeSeriesService } from '../../../services/TimeSeries.service';
import { TimeSeries } from '../../../models/TimeSeries';

@Component({
    selector: 'app-index-timeSeries',
    standalone: false,
    templateUrl: './index.component.html',
    styleUrls: ['./index.component.css']
})
export class IndexTimeSeriesComponent implements OnInit {

    timeSeriess: TimeSeries[] = [];

    constructor(
        private router: Router,
        private service: TimeSeriesService
) {}

    ngOnInit(): void {
        this.getTimeSeriess();
}

    getTimeSeriess(): void {
        this.service.getTimeSeriess().subscribe((res) => {
        this.timeSeriess = res;
    });
}

    deleteTimeSeries(id: any): void {
        this.service.deleteTimeSeries(id)
            .subscribe(() => {
                this.getTimeSeriess();
            });
    }
}