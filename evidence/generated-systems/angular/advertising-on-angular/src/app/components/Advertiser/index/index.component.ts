
import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { AdvertiserService } from '../../../services/Advertiser.service';
import { Advertiser } from '../../../models/Advertiser';

@Component({
    selector: 'app-index-advertiser',
    standalone: false,
    templateUrl: './index.component.html',
    styleUrls: ['./index.component.css']
})
export class IndexAdvertiserComponent implements OnInit {

    advertisers: Advertiser[] = [];

    constructor(
        private router: Router,
        private service: AdvertiserService
) {}

    ngOnInit(): void {
        this.getAdvertisers();
}

    getAdvertisers(): void {
        this.service.getAdvertisers().subscribe((res) => {
        this.advertisers = res;
    });
}

    deleteAdvertiser(id: any): void {
        this.service.deleteAdvertiser(id)
            .subscribe(() => {
                this.getAdvertisers();
            });
    }
}