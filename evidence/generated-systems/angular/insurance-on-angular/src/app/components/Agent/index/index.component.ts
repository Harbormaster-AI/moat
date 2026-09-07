
import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { AgentService } from '../../../services/Agent.service';
import { Agent } from '../../../models/Agent';

@Component({
    selector: 'app-index-agent',
    standalone: false,
    templateUrl: './index.component.html',
    styleUrls: ['./index.component.css']
})
export class IndexAgentComponent implements OnInit {

    agents: Agent[] = [];

    constructor(
        private router: Router,
        private service: AgentService
) {}

    ngOnInit(): void {
        this.getAgents();
}

    getAgents(): void {
        this.service.getAgents().subscribe((res) => {
        this.agents = res;
    });
}

    deleteAgent(id: any): void {
        this.service.deleteAgent(id)
            .subscribe(() => {
                this.getAgents();
            });
    }
}