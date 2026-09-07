
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { Router } from '@angular/router';
import { IndexCarePlanComponent } from './index.component';
import { CarePlanService } from '../../../services/CarePlan.service';

describe('IndexCarePlanComponent', () => {
  let component: IndexCarePlanComponent;
  let fixture: ComponentFixture<IndexCarePlanComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [
        IndexCarePlanComponent
      ],
      providers: [
        CarePlanService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate'),
            navigateByUrl: jasmine.createSpy('navigateByUrl')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(IndexCarePlanComponent);
    component = fixture.componentInstance;

    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});