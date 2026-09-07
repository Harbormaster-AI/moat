
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { Router } from '@angular/router';
import { IndexFeeScheduleComponent } from './index.component';
import { FeeScheduleService } from '../../../services/FeeSchedule.service';

describe('IndexFeeScheduleComponent', () => {
  let component: IndexFeeScheduleComponent;
  let fixture: ComponentFixture<IndexFeeScheduleComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [
        IndexFeeScheduleComponent
      ],
      providers: [
        FeeScheduleService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate'),
            navigateByUrl: jasmine.createSpy('navigateByUrl')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(IndexFeeScheduleComponent);
    component = fixture.componentInstance;

    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});