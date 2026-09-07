
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { Router } from '@angular/router';
import { IndexWorkScheduleComponent } from './index.component';
import { WorkScheduleService } from '../../../services/WorkSchedule.service';

describe('IndexWorkScheduleComponent', () => {
  let component: IndexWorkScheduleComponent;
  let fixture: ComponentFixture<IndexWorkScheduleComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [
        IndexWorkScheduleComponent
      ],
      providers: [
        WorkScheduleService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate'),
            navigateByUrl: jasmine.createSpy('navigateByUrl')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(IndexWorkScheduleComponent);
    component = fixture.componentInstance;

    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});