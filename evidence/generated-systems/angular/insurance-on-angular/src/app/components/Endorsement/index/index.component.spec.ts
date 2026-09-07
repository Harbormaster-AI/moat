
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { Router } from '@angular/router';
import { IndexEndorsementComponent } from './index.component';
import { EndorsementService } from '../../../services/Endorsement.service';

describe('IndexEndorsementComponent', () => {
  let component: IndexEndorsementComponent;
  let fixture: ComponentFixture<IndexEndorsementComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [
        IndexEndorsementComponent
      ],
      providers: [
        EndorsementService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate'),
            navigateByUrl: jasmine.createSpy('navigateByUrl')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(IndexEndorsementComponent);
    component = fixture.componentInstance;

    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});