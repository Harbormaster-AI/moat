
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { Router } from '@angular/router';
import { IndexInspectionPlanComponent } from './index.component';
import { InspectionPlanService } from '../../../services/InspectionPlan.service';

describe('IndexInspectionPlanComponent', () => {
  let component: IndexInspectionPlanComponent;
  let fixture: ComponentFixture<IndexInspectionPlanComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [
        IndexInspectionPlanComponent
      ],
      providers: [
        InspectionPlanService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate'),
            navigateByUrl: jasmine.createSpy('navigateByUrl')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(IndexInspectionPlanComponent);
    component = fixture.componentInstance;

    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});