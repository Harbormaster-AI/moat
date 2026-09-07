
import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { ConversionEventService } from '../../../services/ConversionEvent.service';
import { ConversionEvent } from '../../../models/ConversionEvent';

@Component({
    selector: 'app-index-conversionEvent',
    standalone: false,
    templateUrl: './index.component.html',
    styleUrls: ['./index.component.css']
})
export class IndexConversionEventComponent implements OnInit {

    conversionEvents: ConversionEvent[] = [];

    constructor(
        private router: Router,
        private service: ConversionEventService
) {}

    ngOnInit(): void {
        this.getConversionEvents();
}

    getConversionEvents(): void {
        this.service.getConversionEvents().subscribe((res) => {
        this.conversionEvents = res;
    });
}

    deleteConversionEvent(id: any): void {
        this.service.deleteConversionEvent(id)
            .subscribe(() => {
                this.getConversionEvents();
            });
    }
}