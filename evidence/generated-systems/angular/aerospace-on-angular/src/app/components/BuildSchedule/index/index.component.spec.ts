
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { Router } from '@angular/router';
import { IndexBuildScheduleComponent } from './index.component';
import { BuildScheduleService } from '../../../services/BuildSchedule.service';

describe('IndexBuildScheduleComponent', () => {
  let component: IndexBuildScheduleComponent;
  let fixture: ComponentFixture<IndexBuildScheduleComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [
        IndexBuildScheduleComponent
      ],
      providers: [
        BuildScheduleService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate'),
            navigateByUrl: jasmine.createSpy('navigateByUrl')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(IndexBuildScheduleComponent);
    component = fixture.componentInstance;

    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});