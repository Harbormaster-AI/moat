
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { Router } from '@angular/router';
import { IndexJobFamilyComponent } from './index.component';
import { JobFamilyService } from '../../../services/JobFamily.service';

describe('IndexJobFamilyComponent', () => {
  let component: IndexJobFamilyComponent;
  let fixture: ComponentFixture<IndexJobFamilyComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [
        IndexJobFamilyComponent
      ],
      providers: [
        JobFamilyService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate'),
            navigateByUrl: jasmine.createSpy('navigateByUrl')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(IndexJobFamilyComponent);
    component = fixture.componentInstance;

    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});