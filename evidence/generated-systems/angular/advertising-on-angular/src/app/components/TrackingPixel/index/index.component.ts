
import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { TrackingPixelService } from '../../../services/TrackingPixel.service';
import { TrackingPixel } from '../../../models/TrackingPixel';

@Component({
    selector: 'app-index-trackingPixel',
    standalone: false,
    templateUrl: './index.component.html',
    styleUrls: ['./index.component.css']
})
export class IndexTrackingPixelComponent implements OnInit {

    trackingPixels: TrackingPixel[] = [];

    constructor(
        private router: Router,
        private service: TrackingPixelService
) {}

    ngOnInit(): void {
        this.getTrackingPixels();
}

    getTrackingPixels(): void {
        this.service.getTrackingPixels().subscribe((res) => {
        this.trackingPixels = res;
    });
}

    deleteTrackingPixel(id: any): void {
        this.service.deleteTrackingPixel(id)
            .subscribe(() => {
                this.getTrackingPixels();
            });
    }
}