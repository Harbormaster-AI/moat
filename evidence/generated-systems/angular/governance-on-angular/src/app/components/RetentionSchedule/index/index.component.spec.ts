
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { Router } from '@angular/router';
import { IndexRetentionScheduleComponent } from './index.component';
import { RetentionScheduleService } from '../../../services/RetentionSchedule.service';

describe('IndexRetentionScheduleComponent', () => {
  let component: IndexRetentionScheduleComponent;
  let fixture: ComponentFixture<IndexRetentionScheduleComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [
        IndexRetentionScheduleComponent
      ],
      providers: [
        RetentionScheduleService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate'),
            navigateByUrl: jasmine.createSpy('navigateByUrl')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(IndexRetentionScheduleComponent);
    component = fixture.componentInstance;

    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});