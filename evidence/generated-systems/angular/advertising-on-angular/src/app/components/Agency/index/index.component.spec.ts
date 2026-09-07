
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { Router } from '@angular/router';
import { IndexAgencyComponent } from './index.component';
import { AgencyService } from '../../../services/Agency.service';

describe('IndexAgencyComponent', () => {
  let component: IndexAgencyComponent;
  let fixture: ComponentFixture<IndexAgencyComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [
        IndexAgencyComponent
      ],
      providers: [
        AgencyService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate'),
            navigateByUrl: jasmine.createSpy('navigateByUrl')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(IndexAgencyComponent);
    component = fixture.componentInstance;

    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});