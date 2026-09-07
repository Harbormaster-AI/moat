
import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { ThirdPartyService } from '../../../services/ThirdParty.service';
import { ThirdParty } from '../../../models/ThirdParty';

@Component({
    selector: 'app-index-thirdParty',
    standalone: false,
    templateUrl: './index.component.html',
    styleUrls: ['./index.component.css']
})
export class IndexThirdPartyComponent implements OnInit {

    thirdPartys: ThirdParty[] = [];

    constructor(
        private router: Router,
        private service: ThirdPartyService
) {}

    ngOnInit(): void {
        this.getThirdPartys();
}

    getThirdPartys(): void {
        this.service.getThirdPartys().subscribe((res) => {
        this.thirdPartys = res;
    });
}

    deleteThirdParty(id: any): void {
        this.service.deleteThirdParty(id)
            .subscribe(() => {
                this.getThirdPartys();
            });
    }
}