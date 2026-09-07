
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { Router } from '@angular/router';
import { IndexLineageNodeComponent } from './index.component';
import { LineageNodeService } from '../../../services/LineageNode.service';

describe('IndexLineageNodeComponent', () => {
  let component: IndexLineageNodeComponent;
  let fixture: ComponentFixture<IndexLineageNodeComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [
        IndexLineageNodeComponent
      ],
      providers: [
        LineageNodeService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate'),
            navigateByUrl: jasmine.createSpy('navigateByUrl')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(IndexLineageNodeComponent);
    component = fixture.componentInstance;

    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});