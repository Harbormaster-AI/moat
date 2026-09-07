
import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { LineageNodeService } from '../../../services/LineageNode.service';
import { LineageNode } from '../../../models/LineageNode';

@Component({
    selector: 'app-index-lineageNode',
    standalone: false,
    templateUrl: './index.component.html',
    styleUrls: ['./index.component.css']
})
export class IndexLineageNodeComponent implements OnInit {

    lineageNodes: LineageNode[] = [];

    constructor(
        private router: Router,
        private service: LineageNodeService
) {}

    ngOnInit(): void {
        this.getLineageNodes();
}

    getLineageNodes(): void {
        this.service.getLineageNodes().subscribe((res) => {
        this.lineageNodes = res;
    });
}

    deleteLineageNode(id: any): void {
        this.service.deleteLineageNode(id)
            .subscribe(() => {
                this.getLineageNodes();
            });
    }
}