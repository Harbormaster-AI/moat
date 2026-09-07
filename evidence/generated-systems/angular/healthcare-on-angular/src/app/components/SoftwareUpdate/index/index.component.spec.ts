
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { Router } from '@angular/router';
import { IndexSoftwareUpdateComponent } from './index.component';
import { SoftwareUpdateService } from '../../../services/SoftwareUpdate.service';

describe('IndexSoftwareUpdateComponent', () => {
  let component: IndexSoftwareUpdateComponent;
  let fixture: ComponentFixture<IndexSoftwareUpdateComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [
        IndexSoftwareUpdateComponent
      ],
      providers: [
        SoftwareUpdateService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate'),
            navigateByUrl: jasmine.createSpy('navigateByUrl')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(IndexSoftwareUpdateComponent);
    component = fixture.componentInstance;

    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});