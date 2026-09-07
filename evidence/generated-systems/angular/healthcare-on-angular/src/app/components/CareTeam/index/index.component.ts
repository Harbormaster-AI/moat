
import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { CareTeamService } from '../../../services/CareTeam.service';
import { CareTeam } from '../../../models/CareTeam';

@Component({
    selector: 'app-index-careTeam',
    standalone: false,
    templateUrl: './index.component.html',
    styleUrls: ['./index.component.css']
})
export class IndexCareTeamComponent implements OnInit {

    careTeams: CareTeam[] = [];

    constructor(
        private router: Router,
        private service: CareTeamService
) {}

    ngOnInit(): void {
        this.getCareTeams();
}

    getCareTeams(): void {
        this.service.getCareTeams().subscribe((res) => {
        this.careTeams = res;
    });
}

    deleteCareTeam(id: any): void {
        this.service.deleteCareTeam(id)
            .subscribe(() => {
                this.getCareTeams();
            });
    }
}