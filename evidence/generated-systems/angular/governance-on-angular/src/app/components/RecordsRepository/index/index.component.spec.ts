
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { Router } from '@angular/router';
import { IndexRecordsRepositoryComponent } from './index.component';
import { RecordsRepositoryService } from '../../../services/RecordsRepository.service';

describe('IndexRecordsRepositoryComponent', () => {
  let component: IndexRecordsRepositoryComponent;
  let fixture: ComponentFixture<IndexRecordsRepositoryComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [
        IndexRecordsRepositoryComponent
      ],
      providers: [
        RecordsRepositoryService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate'),
            navigateByUrl: jasmine.createSpy('navigateByUrl')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(IndexRecordsRepositoryComponent);
    component = fixture.componentInstance;

    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});