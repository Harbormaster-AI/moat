
import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { TeamService } from '../../../services/Team.service';
import { Team } from '../../../models/Team';

@Component({
    selector: 'app-index-team',
    standalone: false,
    templateUrl: './index.component.html',
    styleUrls: ['./index.component.css']
})
export class IndexTeamComponent implements OnInit {

    teams: Team[] = [];

    constructor(
        private router: Router,
        private service: TeamService
) {}

    ngOnInit(): void {
        this.getTeams();
}

    getTeams(): void {
        this.service.getTeams().subscribe((res) => {
        this.teams = res;
    });
}

    deleteTeam(id: any): void {
        this.service.deleteTeam(id)
            .subscribe(() => {
                this.getTeams();
            });
    }
}