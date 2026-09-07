
import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { EndorsementService } from '../../../services/Endorsement.service';
import { Endorsement } from '../../../models/Endorsement';

@Component({
    selector: 'app-index-endorsement',
    standalone: false,
    templateUrl: './index.component.html',
    styleUrls: ['./index.component.css']
})
export class IndexEndorsementComponent implements OnInit {

    endorsements: Endorsement[] = [];

    constructor(
        private router: Router,
        private service: EndorsementService
) {}

    ngOnInit(): void {
        this.getEndorsements();
}

    getEndorsements(): void {
        this.service.getEndorsements().subscribe((res) => {
        this.endorsements = res;
    });
}

    deleteEndorsement(id: any): void {
        this.service.deleteEndorsement(id)
            .subscribe(() => {
                this.getEndorsements();
            });
    }
}