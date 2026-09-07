
import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { AudienceSegmentService } from '../../../services/AudienceSegment.service';
import { AudienceSegment } from '../../../models/AudienceSegment';

@Component({
    selector: 'app-index-audienceSegment',
    standalone: false,
    templateUrl: './index.component.html',
    styleUrls: ['./index.component.css']
})
export class IndexAudienceSegmentComponent implements OnInit {

    audienceSegments: AudienceSegment[] = [];

    constructor(
        private router: Router,
        private service: AudienceSegmentService
) {}

    ngOnInit(): void {
        this.getAudienceSegments();
}

    getAudienceSegments(): void {
        this.service.getAudienceSegments().subscribe((res) => {
        this.audienceSegments = res;
    });
}

    deleteAudienceSegment(id: any): void {
        this.service.deleteAudienceSegment(id)
            .subscribe(() => {
                this.getAudienceSegments();
            });
    }
}