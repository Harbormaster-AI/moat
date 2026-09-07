
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { Router } from '@angular/router';
import { IndexProductionScheduleComponent } from './index.component';
import { ProductionScheduleService } from '../../../services/ProductionSchedule.service';

describe('IndexProductionScheduleComponent', () => {
  let component: IndexProductionScheduleComponent;
  let fixture: ComponentFixture<IndexProductionScheduleComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [
        IndexProductionScheduleComponent
      ],
      providers: [
        ProductionScheduleService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate'),
            navigateByUrl: jasmine.createSpy('navigateByUrl')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(IndexProductionScheduleComponent);
    component = fixture.componentInstance;

    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});