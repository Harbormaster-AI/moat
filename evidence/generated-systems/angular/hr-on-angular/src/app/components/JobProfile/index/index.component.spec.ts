
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { Router } from '@angular/router';
import { IndexJobProfileComponent } from './index.component';
import { JobProfileService } from '../../../services/JobProfile.service';

describe('IndexJobProfileComponent', () => {
  let component: IndexJobProfileComponent;
  let fixture: ComponentFixture<IndexJobProfileComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [
        IndexJobProfileComponent
      ],
      providers: [
        JobProfileService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate'),
            navigateByUrl: jasmine.createSpy('navigateByUrl')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(IndexJobProfileComponent);
    component = fixture.componentInstance;

    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});