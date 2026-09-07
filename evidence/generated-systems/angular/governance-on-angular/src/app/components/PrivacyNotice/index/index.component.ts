
import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { PrivacyNoticeService } from '../../../services/PrivacyNotice.service';
import { PrivacyNotice } from '../../../models/PrivacyNotice';

@Component({
    selector: 'app-index-privacyNotice',
    standalone: false,
    templateUrl: './index.component.html',
    styleUrls: ['./index.component.css']
})
export class IndexPrivacyNoticeComponent implements OnInit {

    privacyNotices: PrivacyNotice[] = [];

    constructor(
        private router: Router,
        private service: PrivacyNoticeService
) {}

    ngOnInit(): void {
        this.getPrivacyNotices();
}

    getPrivacyNotices(): void {
        this.service.getPrivacyNotices().subscribe((res) => {
        this.privacyNotices = res;
    });
}

    deletePrivacyNotice(id: any): void {
        this.service.deletePrivacyNotice(id)
            .subscribe(() => {
                this.getPrivacyNotices();
            });
    }
}