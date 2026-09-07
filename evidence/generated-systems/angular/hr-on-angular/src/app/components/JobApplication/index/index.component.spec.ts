
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { Router } from '@angular/router';
import { IndexJobApplicationComponent } from './index.component';
import { JobApplicationService } from '../../../services/JobApplication.service';

describe('IndexJobApplicationComponent', () => {
  let component: IndexJobApplicationComponent;
  let fixture: ComponentFixture<IndexJobApplicationComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [
        IndexJobApplicationComponent
      ],
      providers: [
        JobApplicationService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate'),
            navigateByUrl: jasmine.createSpy('navigateByUrl')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(IndexJobApplicationComponent);
    component = fixture.componentInstance;

    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});