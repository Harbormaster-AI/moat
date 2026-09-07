
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { Router } from '@angular/router';
import { IndexScheduleExceptionComponent } from './index.component';
import { ScheduleExceptionService } from '../../../services/ScheduleException.service';

describe('IndexScheduleExceptionComponent', () => {
  let component: IndexScheduleExceptionComponent;
  let fixture: ComponentFixture<IndexScheduleExceptionComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [
        IndexScheduleExceptionComponent
      ],
      providers: [
        ScheduleExceptionService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate'),
            navigateByUrl: jasmine.createSpy('navigateByUrl')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(IndexScheduleExceptionComponent);
    component = fixture.componentInstance;

    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});